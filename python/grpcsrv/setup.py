# -*- coding:utf-8 -*-

import codecs
import os

from setuptools import find_packages, setup


def read(fname):
    return codecs.open(os.path.join(
        os.path.dirname(__file__), fname)).read().strip()


def read_install_requires():
    reqs = [
            'numpy>=1.16.6'
            ]
    return reqs


setup(name='grpcsrv',
      version=read('VERSION.txt'),
      description='grpc service example',
      long_description_content_type='text/markdown',
      long_description=read('README.md'),
      url='https://github.com/iminders/codelabs/tree/master/python/grpcsrv',
      author='iminders',
      author_email='liuwen.w@qq.com',
      packages=find_packages(),
      license="MIT License",
      zip_safe=False,
      classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        ],
      python_requires='>=3',
      platforms='any',
      package_data={'': ['LICENSE', 'README.md']},
      install_requires=read_install_requires(),
      include_package_data=True
      )
