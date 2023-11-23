#ifndef DRIVERSDK_EVENT_H
#define DRIVERSDK_EVENT_H

#include <string>

namespace DRIVERSDK {
    class Event {
    private:
        std::string type;
        void *payload;

    public:
        Event(const std::string &eventType, void *eventPayload);

        const std::string &getType() const;

        void setType(const std::string &eventType);

        const void *getPayload() const;

        void setPayload(void *eventPayload);
    };
}

#endif //DRIVERSDK_EVENT_H
