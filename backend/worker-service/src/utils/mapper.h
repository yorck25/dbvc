#include <pqxx/pqxx>

class Mapper {
public:
    static Project mapProject(const pqxx::row &row);
};