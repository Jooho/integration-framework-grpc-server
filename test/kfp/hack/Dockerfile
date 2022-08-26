FROM python
RUN apt-get update;apt-get install vim -y
RUN pip install --upgrade pip;pip install kfp-tekton
RUN mkdir -p /home/app/src

WORKDIR /home
CMD python /home/app/src/deploy-model.py
