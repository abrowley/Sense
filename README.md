# Sense

## Overview
Sense is a server which receives JSON data from sensor in the form of simple http POST requests, stores them in a NoSQL database and notifies clients connected via the websocket protocol of the update. 

It is intended to be used for internet of things applications.

It is written in golang and uses mongodb
