#!/bin/sh

# echo $PHOTON_OPTS

java $JAVA_OPTS -jar photon-0.3.5.jar -nominatim-import $PHOTON_OPTS

java $JAVA_OPTS -jar photon-0.3.5.jar $PHOTON_OPTS