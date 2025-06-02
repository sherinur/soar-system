CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" VARCHAR(128) UNIQUE NOT NULL,
  "password_hash" VARCHAR(255) NOT NULL,
  "first_name" VARCHAR(64) NOT NULL,
  "last_name" VARCHAR(64) NOT NULL,
  "role" user_role NOT NULL,
  "organization_id" INTEGER NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "FK_User_Organization"
    FOREIGN KEY ("organization_id")
    REFERENCES "Organization"("id")
    ON DELETE CASCADE
);