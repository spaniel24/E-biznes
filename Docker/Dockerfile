FROM openjdk:8-alpine

RUN apk add --no-cache curl bash

RUN mkdir -p /usr/share/gradle \
  && echo "Downlaoding gradle..." \
  && curl -fsSL -o /tmp/gradle.zip https://services.gradle.org/distributions/gradle-7.4.2-bin.zip \
  \
  && echo "Extracting gradle..." \
  && unzip -d /usr/share/gradle /tmp/gradle.zip \
  \
  && echo "Cleaning up..." \
  && rm -f /tmp/gradle.zip \
  && ln -s /usr/share/gradle/gradle-7.4.2 /usr/bin/gradle \
  \
  && echo "Gradle installed!"

ENV GRADLE_VERSION 7.4.2
ENV GRADLE_HOME /usr/bin/gradle
ENV GRADLE_USER_HOME /cache
ENV PATH $PATH:$GRADLE_HOME/bin

VOLUME docker_test_volume

EXPOSE 8000