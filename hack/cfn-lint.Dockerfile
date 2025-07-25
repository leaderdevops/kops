FROM python:3.14.0rc1-alpine
ARG CFNLINT_VERSION
RUN pip install "cfn-lint==${CFNLINT_VERSION}" pydot
ENTRYPOINT ["cfn-lint"]