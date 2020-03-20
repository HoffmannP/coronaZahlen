#!/bin/bash

name="js/script.js"
echo "put ichplatz/${name} ichplatz/coronaZahlen/${name}" | ftp -v ichplatz.de
