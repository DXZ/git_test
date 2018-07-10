import sys
sys.path.insert(0, "./")
from tensorflow_serving_client.protos import predict_pb2, prediction_service_pb2,prediction_service_pb2_grpc

# from grpc.beta import implementations
import grpccv2
import tensorflow as tf
import grpc
from tensorflow.python.framework import dtypes
import time
import numpy as np
import cv2
#注意，如果在windows下测试，文件名可能需要写成：im_name = r"测试文件目录\文件名"
im_name = "/data/seven/dl/data/small/data/20180111_27701_typeB.jpg"
# if __name__ == '__main__':
#文件读取和处理
im = cv2.imread(im_name)
re_im = cv2.resize(im, (224, 224), interpolation=cv2.INTER_CUBIC)
#记个时
start_time = time.time()
#建立连接
host = "localhost"
port = 9000
# channel = implementations.insecure_channel("localhost", 9000)
channel = grpc.insecure_channel('%s:%s' % (host, port))
# stub = prediction_service_pb2.beta_create_PredictionService_stub(channel)
stub = prediction_service_pb2_grpc.PredictionServiceStub(channel)
request = predict_pb2.PredictRequest()
#这里由保存和运行时定义，第一个是运行时配置的模型名，第二个是保存时输入的方法名
request.model_spec.name = "mnist"
#入参参照入参定义
request.model_spec.signature_name = 'predict'


width = 286
height = 32
raw_files = []

f = open(im_name,'rb')
raw_files.append(f)
n = len(raw_files)
X = np.ones((n, width, height, 1), dtype=np.uint8) * 127
for i in range(n):
    raw_file = raw_files[i]
    raw_file.seek(0)
    file_bytes = np.asarray(bytearray(raw_file.read()), dtype=np.uint8)
    img = cv2.imdecode(file_bytes, 0)
    resize_num = float(height)/float(img.shape[0])
    if resize_num < 1:
        img = cv2.resize(img,(int(img.shape[1]*resize_num),int(img.shape[0]*resize_num)))
    X[i, :img.shape[1], :img.shape[0], 0] = img.transpose(1, 0)

x = X.astype(np.float32)
request.inputs["images"].ParseFromString(tf.contrib.util.make_tensor_proto(x).SerializeToString())
    #第二个参数是最大等待时间，因为这里是block模式response = stub.Predict(request, 10.0)访问的

response = stub.Predict(request, 10)
nd_array = tf.contrib.util.make_ndarray(response.outputs["scores"])

characters = '=&0123456789.+-×÷<>(){}[]@%#| '

pred = nd_array[0]


y_sorted = pred.argsort()[:,::-1]
# print y_sorted[:,:10]
res = ''
last_char = ''
for index,value in enumerate(y_sorted):
    new_char = characters[value[0]]
    if new_char != last_char :
        res += new_char
    last_char = new_char

print(res)


results = {}
for key in response.outputs:
    tensor_proto = response.outputs[key]
    print("aaaaa",key,tensor_proto)
    nd_array = tf.contrib.util.make_ndarray(tensor_proto)
    results[key] = nd_array
    print("cost %ss to predict: " % (time.time() - start_time))
    print(results)

