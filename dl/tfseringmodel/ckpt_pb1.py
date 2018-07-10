# coding=UTF-8
import tensorflow as tf
import os.path
import argparse
from tensorflow.python.framework import graph_util
from tensorflow.python.saved_model.signature_def_utils_impl import  build_signature_def, predict_signature_def



from lib.networks.factory import get_network
from lib.fast_rcnn.config import cfg,cfg_from_file


MODEL_DIR = "./tfpb"
MODEL_NAME = "frozen_model.pb"

if not tf.gfile.Exists(MODEL_DIR): #创建目录
    tf.gfile.MakeDirs(MODEL_DIR)


def export_saved_model(version, path, sess=None):

    net = get_network("VGGnet_test")
    im = 128 * np.ones((300, 300, 3), dtype=np.uint8)


    
    



    tf.app.flags.DEFINE_integer('version', version, 'version number of the model.')
    tf.app.flags.DEFINE_string('work_dir', path, 'your older model  directory.')
    tf.app.flags.DEFINE_string('model_dir', '/tmp/model_name', 'saved model directory')
    FLAGS = tf.app.flags.FLAGS

    # you can give the session and export your model immediately after training 
    if not sess: 
        saver = tf.train.import_meta_graph(os.path.join(path, 'xxx.ckpt.meta'))
        saver.restore(sess, tf.train.latest_checkpoint(path))

    export_path = os.path.join(
        tf.compat.as_bytes(FLAGS.model_dir),
        tf.compat.as_bytes(str(FLAGS.version)))
    builder = tf.saved_model.builder.SavedModelBuilder(export_path)

    # define the signature def map here
    # ...


    os.environ["CUDA_VISIBLE_DEVICES"] = "7"
    config = tf.ConfigProto()
    config.gpu_options.per_process_gpu_memory_fraction = 0.2
    graph = tf.Graph().as_default()
    # with tf.Graph().as_default() as graph:
    # sess = tf.Session()
    # MODEL_FILE = 'model/tensorflow_inception_graph.pb'
    # BOTTLENECK_TENSOR_NAME = 'pool_3/_reshape:0'  # tensor name for the bottleneck of inception-v3 model
    # # JPEG_DATA_TENSOR_NAME = 'DecodeJpeg/contents:0'
    # JPEG_DATA_TENSOR_NAME = 'DecodeJpeg:0'  # image tensor
    # f = tf.gfile.FastGFile(MODEL_FILE, 'rb')
    # # import initial pb model
    # graph_def = tf.GraphDef()
    # graph_def.ParseFromString(f.read())
    # bottleneck_tensor, jpeg_data_tensor = tf.import_graph_def(graph_def,return_elements=[BOTTLENECK_TENSOR_NAME, JPEG_DATA_TENSOR_NAME])

    prediction_signature = predict_signature_def(inputs={'image': graph.get_operation_by_name('BottleneckInputPlaceholder').outputs[0]},
                                  outputs={'scores': graph.get_operation_by_name('evaluation/ArgMax').outputs[0]})

    # legacy_init_op = tf.group(tf.tables_initializer(), name='legacy_init_op')
    builder.add_meta_graph_and_variables(
        sess=sess,
        tags=[tf.saved_model.tag_constants.SERVING],
        signature_def_map={
            'predict':
                prediction_signature
        },
    )

    builder.save()
    print('Export SavedModel!')


if __name__ == '__main__':
    # parser = argparse.ArgumentParser()
    # parser.add_argument("model_folder", type=str, help="input ckpt model dir") #命令行解析，help是提示符，type是输入的类型，
    # # 这里运行程序时需要带上模型ckpt的路径，不然会报 error: too few arguments
    # aggs = parser.parse_args()
    freeze_graph("/data/seven/dl/tfserving/testmodel/ctpnModel/release")



import tensorflow as tf
import os.path
import argparse
from tensorflow.python.framework import graph_util
from tensorflow.python.saved_model.signature_def_utils_impl import  build_signature_def, predict_signature_def


os.environ["CUDA_VISIBLE_DEVICES"] = "7"

sess = tf.Session()
path = "/data/knowbox/seven/googlenet/model/checkpoints"
saver = tf.train.import_meta_graph(os.path.join(path, 'model-23900.meta'))
saver.restore(sess, tf.train.latest_checkpoint(path))

export_path = "/data/knowbox/wangth/workspace/filter/Inception_v3_classifier/model/output/3"
builder = tf.saved_model.builder.SavedModelBuilder(export_path)

# define the signature def map here
# ...

graph = tf.Graph().as_default()
# sess = tf.Session()
MODEL_FILE = '/data/knowbox/wangth/workspace/filter/Inception_v3_classifier/model/output/1/saved_model.pb'
BOTTLENECK_TENSOR_NAME = 'pool_3/_reshape:0'  # tensor name for the bottleneck of inception-v3 model
JPEG_DATA_TENSOR_NAME = 'DecodeJpeg:0'  # image tensor
f = tf.gfile.FastGFile(MODEL_FILE, 'rb')
# # import initial pb model
graph_def = tf.GraphDef()
graph_def.ParseFromString(f.read())
tf.import_graph_def(graph_def, name='')
# bottleneck_tensor, jpeg_data_tensor = tf.import_graph_def(graph_def,return_elements=[BOTTLENECK_TENSOR_NAME, JPEG_DATA_TENSOR_NAME])
out_tensor = graph_def.get_tensor_by_name('final_result:0')
input_tensor = graph_def.get_tensor_by_name('DecodeJpeg:0')

# prediction_signature = predict_signature_def(inputs={'image': graph.get_operation_by_name('BottleneckInputPlaceholder').outputs[0]},
#                               outputs={'scores': graph.get_operation_by_name('evaluation/ArgMax').outputs[0]})

prediction_signature = predict_signature_def(inputs={'image':input_tensor},
                              outputs={'result':out_tensor})

# legacy_init_op = tf.group(tf.tables_initializer(), name='legacy_init_op')
builder.add_meta_graph_and_variables(
    sess=sess,
    tags=[tf.saved_model.tag_constants.SERVING],
    signature_def_map={'predict':prediction_signature},
)

builder.save()
print('Export SavedModel!')
