# coding=utf-8

from __future__ import print_function
from grpc.beta import implementations
import tensorflow as tf

from tensorflow_serving.apis import predict_pb2
from tensorflow_serving.apis import prediction_service_pb2
import numpy as np

tf.app.flags.DEFINE_string('server', 'localhost:9000',
                           'PredictionService host:port')
tf.app.flags.DEFINE_string('vocie', '', 'path to voice in wav format')
FLAGS = tf.app.flags.FLAGS

def get_melgram(path):
    melgram = .... # 这里省略
    return melgram

def main(_):
    host, port = FLAGS.server.split(':')
    channel = implementations.insecure_channel(host, int(port))
    stub = prediction_service_pb2.beta_create_PredictionService_stub(channel)
    # Send request
 
    # See prediction_service.proto for gRPC request/response details.
    data = get_melgram("T_1000001.wav")
    data = data.astype(np.float32)

    request = predict_pb2.PredictRequest()
    request.model_spec.name = 'voice' # 这个name跟tensorflow_model_server  --model_name="voice" 对应
    request.model_spec.signature_name = 'voice_classification' # 这个signature_name  跟signature_def_map 对应
    request.inputs['voice'].CopyFrom(
          tf.contrib.util.make_tensor_proto(data, shape=[1, 1, 96, 89])) # shape跟 keras的model.input类型对应
    result = stub.Predict(request, 10.0)  # 10 secs timeout
    print(result)


if __name__ == '__main__':
  tf.app.run()
