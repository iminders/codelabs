# -*- coding:utf-8 -*-

import os

import tensorflow as tf
from tensorflow import keras

from tensorflowsrv.common.logger import logger

MODEL_DIR = "testdata/faishon"


def main():

    logger.info('TensorFlow version: {}'.format(tf.__version__))

    fashion_mnist = keras.datasets.fashion_mnist
    (train_images, train_labels), (
        test_images, test_labels) = fashion_mnist.load_data()

    # scale the values to 0.0 to 1.0
    train_images = train_images / 255.0
    test_images = test_images / 255.0

    # reshape for feeding into the model
    train_images = train_images.reshape(train_images.shape[0], 28, 28, 1)
    test_images = test_images.reshape(test_images.shape[0], 28, 28, 1)

    # class_names = ['T-shirt/top', 'Trouser', 'Pullover', 'Dress', 'Coat',
    #                'Sandal', 'Shirt', 'Sneaker', 'Bag', 'Ankle boot']

    logger.info('\ntrain_images.shape: {}, of {}'.format(
        train_images.shape, train_images.dtype))
    logger.info('test_images.shape: {}, of {}'.format(
        test_images.shape, test_images.dtype))

    model = keras.Sequential([
        keras.layers.Conv2D(input_shape=(28, 28, 1), filters=8, kernel_size=3,
                            strides=2, activation='relu', name='Conv1'),
        keras.layers.Flatten(),
        keras.layers.Dense(10, activation=tf.nn.softmax, name='Softmax')
    ])
    model.summary()

    epochs = 5

    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['accuracy'])
    model.fit(train_images, train_labels, epochs=epochs)

    test_loss, test_acc = model.evaluate(test_images, test_labels)
    logger.info('\nTest accuracy: {}'.format(test_acc))

    version = 1
    export_path = os.path.join(MODEL_DIR, str(version))
    logger.info('export_path = {}\n'.format(export_path))

    tf.keras.models.save_model(
        model,
        export_path,
        overwrite=True,
        include_optimizer=True,
        save_format=None,
        signatures=None,
        options=None
    )

    logger.info('\nSaved model:')


if __name__ == '__main__':
    main()
