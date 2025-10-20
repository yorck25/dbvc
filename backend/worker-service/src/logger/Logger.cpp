#include "Logger.h"

static const char* LogTypes[] = {"Warning", "Info", "Error"};

void Logger::print(LogType type, const std::string& msg) {
    std::cout << "[" << LogTypes[static_cast<int>(type)] << "]: " << msg << std::endl;
}

void Logger::print(const std::string& msg) {
    print(LogType::INFO, msg);
}
