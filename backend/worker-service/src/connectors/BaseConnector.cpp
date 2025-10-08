//
// Created by Yorck Dombrowsky on 07.10.25.
//

#include <iostream>

class BaseConnector {
public:
    std::string host;
    int port;
    std::string user;
    std::string password;
    std::string database;

    virtual ~BaseConnector() = default;
    int maxRetries = 5;

    bool connect(){
        bool state = false;

        for(int i = 0; i < this->maxRetries; i++){
            bool res = tryConnect();
            if(res){
                state = true;
                break;
            }
        }

        return state;
    }

private:
    virtual bool tryConnect(){
        std::println("No tryConnect method implemented.");
        return false;
    }
};