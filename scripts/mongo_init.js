print(
  "Start #################################################################"
);

db.createUser({
  user: "admin",
  pwd: "admin",
  roles: [
    {
      role: "readWrite",
      db: "mydb",
    },
  ],
});

db = db.getSiblingDB("users");
db.createUser({
  user: "api",
  pwd: "api",
  roles: [
    {
      role: "readWrite",
      db: "users",
    },
  ],
});
db.createCollection("users");

db = db.getSiblingDB("messages");
db.createUser({
  user: "chat",
  pwd: "chat",
  roles: [
    {
      role: "readWrite",
      db: "messages",
    },
  ],
});
db.createCollection("messages");
db.messages.createIndex({ timestamp: 1 });
