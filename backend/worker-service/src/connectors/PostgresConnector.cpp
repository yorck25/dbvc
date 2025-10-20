#include "PostgresConnector.h"
#include <iostream>

PostgresConnector::PostgresConnector(std::string host, int port, std::string user,
                                     std::string password, std::string database)
        : host(std::move(host)), port(port), user(std::move(user)),
          password(std::move(password)), database(std::move(database)) {}

bool PostgresConnector::connect() {
    for (int i = 0; i < maxRetries; i++) {
        if (tryConnect()) return true;
    }
    return false;
}

bool PostgresConnector::tryConnect() {
    try {
        connection = std::make_unique<pqxx::connection>(
                "dbname=" + database +
                " user=" + user +
                " password=" + password +
                " host=" + host +
                " port=" + std::to_string(port)
        );
        if (!connection->is_open()) return false;
        pqxx::work txn(*connection);
        pqxx::result r = txn.exec("SELECT version()");
        std::cout << "PostgreSQL version: " << r[0][0].c_str() << std::endl;
        txn.commit();
    } catch (const std::exception &e) {
        std::cerr << "Connection error: " << e.what() << std::endl;
        return false;
    }
    return true;
}

std::optional<pqxx::result> PostgresConnector::executeQuery(const std::string &query) {
    if (!connection || !connection->is_open()) return std::nullopt;
    try {
        pqxx::work txn(*connection);
        pqxx::result r = txn.exec(query);
        txn.commit();
        return r;
    } catch (const std::exception &e) {
        std::cerr << "Fetch error: " << e.what() << std::endl;
        return std::nullopt;
    }
}
