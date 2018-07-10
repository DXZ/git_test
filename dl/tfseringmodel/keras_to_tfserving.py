# coding=utf-8


from keras import backend as K
import sys
import os.path as osp
import os
import tensorflow as tf
from keras.models import load_model


K.set_learning_phase(0)  # all new operations will be in test mode from now on


input_fld = sys.path[0]
weight_file = 'model.h5'
output_graph_name = 'tensor_model.pb'

# output_fld = input_fld + '/tensorflow_model/'
# if not os.path.isdir(output_fld):
#     os.mkdir(output_fld)
weight_file_path = osp.join(input_fld, weight_file)
previous_model = load_model(weight_file_path)


print('input is :', previous_model.input.name)
print ('output is:', previous_model.output.name)

sess = K.get_session()

model = previous_model

# # serialize the model and get its weights, for quick re-building
# config = previous_model.get_config()
# weights = previous_model.get_weights()

# # re-build a model where the learning phase is now hard-coded to 0
# from keras.models import model_from_config
# model = model_from_config(config)
# model.set_weights(weights)



from tensorflow_serving.session_bundle import exporter

export_path = "saved_tfsering.pb" # where to save the exported graph
export_version = 1 # version number (integer)

saver = tf.train.Saver(sharded=True)
model_exporter = exporter.Exporter(saver)
signature = exporter.classification_signature(input_tensor=model.input,
                                              scores_tensor=model.output)
model_exporter.init(sess.graph.as_graph_def(),
                    default_graph_signature=signature)
model_exporter.export(export_path, tf.constant(export_version), sess)