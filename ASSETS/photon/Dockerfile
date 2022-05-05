FROM blacktop/elasticsearch:7.10

ADD . /photon

WORKDIR /photon

RUN wget https://github.com/komoot/photon/releases/download/0.3.5/photon-0.3.5.jar

# RUN wget -O - http://download1.graphhopper.com/public/photon-db-latest.tar.bz2 | bzip2 -cd | tar x

EXPOSE 2322

ENV JAVA_OPTS=""

ENV PHOTON_OPTS=""

RUN ["chmod", "+x", "./startup.sh"]

# RUN java $JAVA_OPTS -jar photon-0.3.5.jar -nominatim-import $PHOTON_OPTS

ENTRYPOINT ["sh", "./startup.sh"]

# ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -jar photon-0.3.5.jar $PHOTON_OPTS"]