cmake_minimum_required(VERSION 3.29)
project(worker_service)

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(worker_service
        src/main.cpp
        src/connectors/BaseConnector.cpp
        src/connectors/PostgresConnector.cpp
        src/utils/projects.cpp
        src/utils/mapper.cpp
        src/utils/mapper.cpp
        src/utils/mapper.cpp
        src/logger/Logger.h
        src/utils/mapper.h
        src/utils/projects.h)

# Homebrew libpqxx headers and libraries
set(PQXX_INCLUDE_DIR /opt/homebrew/Cellar/libpqxx/7.10.3/include)
set(PQXX_LIBRARY /opt/homebrew/Cellar/libpqxx/7.10.3/lib/libpqxx.dylib)
set(PQ_LIBRARY /opt/homebrew/opt/libpq/lib/libpq.dylib)

# Include headers
target_include_directories(worker_service PRIVATE ${PQXX_INCLUDE_DIR} /opt/homebrew/opt/libpq/include)

# Link libraries using absolute paths
target_link_libraries(worker_service PRIVATE ${PQXX_LIBRARY} ${PQ_LIBRARY})
