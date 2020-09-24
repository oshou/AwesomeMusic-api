FROM python:3.7-alpine
WORKDIR /
RUN pip3 install \
  'pytest' \
  'pytest-sugar' \
  'tavern'
COPY ./*.yaml /
ENTRYPOINT ["pytest","test_api.tavern.yaml","-vv","--color=yes","--tavern-global-cfg=local.yaml"]
