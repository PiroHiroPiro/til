FROM python:3.7

WORKDIR /usr/src

RUN pip install --upgrade pip \
    && pip install jupyter

COPY . .

CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--allow-root", "--NotebookApp.token="]
