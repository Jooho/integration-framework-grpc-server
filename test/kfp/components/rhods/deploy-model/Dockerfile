ARG BASE_IMAGE_TAG=slim
FROM docker.io/library/python:$BASE_IMAGE_TAG

USER root

RUN mkdir /app
RUN mkdir /app/src
COPY requirements.txt .

RUN pip install --upgrade pip
RUN python3 -m pip install -r \
    requirements.txt --quiet --no-cache-dir \
    && rm -f requirements.txt
COPY ./src /app/src

RUN chown 1000.1000 -R /app

USER 1000
CMD python3 /app/src/deploy-model.py