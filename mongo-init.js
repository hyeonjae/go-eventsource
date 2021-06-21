db.auth('admin', 'admin')

db = db.getSiblingDB('eventsource')

db.createUser({
    user: 'dev',
    pwd: 'eventsource!@#',
    roles: [
        {
            role: 'dbOwner',
            db: 'eventsource',
        },
    ],
})
