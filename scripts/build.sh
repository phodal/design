#!/usr/bin/env bash

cd languages

antlr -Dlanguage=Go -visitor -listener Design.g4 -o design
