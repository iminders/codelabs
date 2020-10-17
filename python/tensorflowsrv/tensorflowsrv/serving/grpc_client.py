# -*- coding:utf-8 -*-

import json
import time

import grpc
# Helper libraries
import matplotlib.pyplot as plt
import numpy as np
import tensorflow as tf
from tensorflow import keras
# from tensorflow.contrib.util import make_tensor_proto
from tensorflow_serving.apis import predict_pb2, prediction_service_pb2_grpc

fashion_mnist = keras.datasets.fashion_mnist
(train_images, train_labels), (test_images, test_labels) = \
    fashion_mnist.load_data()

# scale the values to 0.0 to 1.0
train_images = train_images / 255.0
test_images = test_images / 255.0

# reshape for feeding into the model
train_images = train_images.reshape(train_images.shape[0], 28, 28, 1)
test_images = test_images.reshape(test_images.shape[0], 28, 28, 1)

class_names = ['T-shirt/top', 'Trouser', 'Pullover', 'Dress', 'Coat',
               'Sandal', 'Shirt', 'Sneaker', 'Bag', 'Ankle boot']

print('\ntrain_images.shape: {}, of {}'.format(train_images.shape,
                                               train_images.dtype))
print('test_images.shape: {}, of {}'.format(test_images.shape,
                                            test_images.dtype))

data = json.dumps({"signature_name": "serving_default",
                  "instances": test_images[0:3].tolist()})
print('Data: {} ... {}'.format(data[:50], data[len(data)-52:]))


def show_result(idx, title):
    plt.figure()
    plt.imshow(test_images[idx].reshape(28, 28))
    plt.axis('off')
    plt.title('\n\n{}'.format(title), fontdict={'size': 16})
    plt.show()


def run(host, port, model, signature_name):
    channel = grpc.insecure_channel('{host}:{port}'.format(
        host=host, port=port))
    stub = prediction_service_pb2_grpc.PredictionServiceStub(channel)

    start = time.time()

    # Call classification model to make prediction on the image
    request = predict_pb2.PredictRequest()
    request.model_spec.name = model
    request.model_spec.signature_name = signature_name
    request.inputs['Conv1_input'].CopyFrom(tf.make_tensor_proto(
        test_images[:3], shape=[3, 28, 28, 1], dtype=tf.float32))

#     print(request)

    result = stub.Predict(request, 5.0)

    end = time.time()
    time_diff = end - start
    print("grpc call time:", time_diff)

    # Reference:
    # How to access nested values
    # https://stackoverflow.com/questions/44785847/how-to-retrieve-float-val-from-a-predictresponse-object
#     print(result)
    predictions = tf.make_ndarray(result.outputs["Softmax"])
    # print(predictions)

    for i in range(0, 3):
        title = 'inference:{} (class {}), actually:{} (class {})'.format(
            class_names[np.argmax(predictions[i])], np.argmax(predictions[i]),
            class_names[test_labels[i]], test_labels[i])
        show_result(i, title)


host = "localhost"
port = 8500
model = "fashion"
signature_name = "serving_default"
run(host, port, model, signature_name)
