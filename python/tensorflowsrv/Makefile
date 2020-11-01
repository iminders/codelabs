test:
	docker-compose up

install:
	pip install -e .

install3:
	pip3 install -e . --user

ltest:
	python3 -m pytest -s

deploy:
	docker run -t --rm -p 8501:8501 -p 8500:8500 \
    -v "$(PWD)/tensorflowsrv/serving/testdata/fashion:/models/fashion" \
    -e MODEL_NAME=fashion \
    tensorflow/serving
