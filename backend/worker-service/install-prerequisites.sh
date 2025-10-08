#!/bin/bash

# Install Postgres Client
brew install postgresql libpqxx
export PKG_CONFIG_PATH="$(brew --prefix libpq)/lib/pkgconfig:$(brew --prefix libpqxx)/lib/pkgconfig:$PKG_CONFIG_PATH"
pkg-config --cflags --libs libpqxx
# Expected something like this: -I/usr/include -lpqxx -lpq
# --------------------------------

