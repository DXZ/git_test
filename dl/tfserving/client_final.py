#coding:utf-8

import sys
import grpc
import cv2
from tensorflow.python.framework import dtypes
import time
from io import BytesIO
import base64
import datetime
import numpy as np
import tensorflow as tf
from flask import Flask, request, jsonify, abort
from tensorflow_serving_client.protos import predict_pb2, prediction_service_pb2,prediction_service_pb2_grpc
import logging
import traceback


logging.basicConfig(level=logging.DEBUG,
                    format='%(asctime)s %(filename)s[line:%(lineno)d] %(levelname)s %(message)s',  
                    datefmt='%a, %d %b %Y %H:%M:%S',  
                    filename='./parser.log',  
                    filemode='a')

logging.info("logging config done")

app = Flask(__name__)
characters = '=&0123456789.+-×÷<>(){}[]@%#| '
#characters = '=120345+6}{１-87×０9２÷４３６５８７９.)(．[]＞＜#＝@＠＋－<%（> '
ocr_model = None
width = 286
height = 32
host, port = "127.0.0.1",9000

channel = grpc.insecure_channel('%s:%s' % (host, port))
stub = prediction_service_pb2_grpc.PredictionServiceStub(channel)
'''
grpc client
'''
def get_grpc_result(nparray):
    request = predict_pb2.PredictRequest()
    request.model_spec.name = "mnist"
    request.model_spec.signature_name = 'predict'
    # x = nparray.astype(np.float32)
    request.inputs["images"].ParseFromString(tf.contrib.util.make_tensor_proto(x).SerializeToString())
    response = stub.Predict(request, 10)
    result_array = tf.contrib.util.make_ndarray(response.outputs["scores"])
    return result_array

def get_result(pred,result_length=3):
    global characters
    y_sorted = pred.argsort()[:,::-1]
    # print y_sorted[:,:10]
    res = ''
    last_char = ''
    for index,value in enumerate(y_sorted):
        new_char = characters[value[0]]
        if new_char != last_char :
            res += new_char
        last_char = new_char
    # print res
    return res.replace(' ','')

def predict_json(raw_files,width):
    global height,ocr_model
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
    start = datetime.datetime.now()
    # y_pred = ocr_model.predict(X,batch_size=n)
    y_pred = get_grpc_result(X)
    print("predict:",datetime.datetime.now() - start)
    start = datetime.datetime.now()
    out = []
    for i in range(n):
        out.append(get_result(y_pred[i]))
    print("ctc_decode:",datetime.datetime.now()- start)
    return out


def dict_to_list(old_dict,run_mode):
    id_list = []
    img_list = []
    name_list = []
    for images in old_dict:
        pic = BytesIO()
        id_list.append(images['id'])
        # if run_mode == 'dev':
        #     name_list.append(images['filename'])
        pic.write(base64.b64decode(images['imgbase64']))
        if pic.read() == '':
            continue
        img_list.append(pic)
    return id_list,img_list,name_list


@app.route("/serving/api/v1.0/images", methods=['POST'])
def reqs():
    logging.info("------start------")
    result = {}
    result["code"] = 500
    try:
        run_mode = None
        data = request.get_json(force=True)
        # print(data['maxwidth'])
        # print(data['run_mode'])
        if not data or not 'reqid' in data:
            abort(400)
        if not 'reqtime' in data:
            abort(400)
        if len(data['images']) > 64:
            abort(400)
        if not 'maxwidth' in data:
            width = 286
        else:
            width = data['maxwidth']
            if width > 1024:
                width = 1024
        if 'run_mode' in data:
            run_mode = data['run_mode']
        id_list,img_list,name_list  = dict_to_list(data['images'],run_mode)
        get_result = predict_json(img_list,width)
        row_result = []
        for i in range(len(img_list)):
            row = {}
            row['id'] = id_list[i]
            row['data'] = [
                    {
                    "pos":[0,0],
                    "detail": [
                        {"content":get_result[i],"pro":1},
                        {"content":get_result[i],"pro":1},
                        {"content":get_result[i],"pro":1}
                        ]
                    }
                    ]
            row_result.append(row)
        result = {
            'code' : 0,
            'retid' : data['reqid'],
            'rettime': datetime.datetime.now().strftime('%Y%m%d%H%M%S'),
            'result': row_result
        }
    except Exception as e:
        exstr = traceback.format_exc()
        logging.error(exstr)
        result["msg"] = exstr
    return jsonify(result)



def run(host='127.0.0.1', port=6000, gpus="0", mem_fra=0.21, model_path="model.h5"):
    """Run a WSGI server using gevent."""
    # WSGIServer((host, port), app).serve_forever()
    app.run(debug=True,host=host,port=port)



if __name__=='__main__':
    run(host='0.0.0.0', port=6000,gpus="0", mem_fra=0.14, model_path="/data/model/parser/release/model.h5")