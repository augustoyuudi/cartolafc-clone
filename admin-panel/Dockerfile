FROM python:3.10.2-slim

RUN useradd -ms /bin/bash cartola

RUN pip install pipenv

USER cartola

WORKDIR /home/cartola/app

ENV PIPENV_VENV_IN_PROJECT=True

# CMD ["./.docker/start.dev.sh"]

CMD ["tail", "-f", "/dev/null"]