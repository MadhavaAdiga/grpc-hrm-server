ALTER TABLE IF EXISTS roles DROP CONSTRAINT IF EXISTS "roles_organization_fkey";
ALTER TABLE IF EXISTS employees DROP CONSTRAINT IF EXISTS "employees_organization_fkey";
ALTER TABLE IF EXISTS employees DROP CONSTRAINT IF EXISTS "employees_user_fkey";

DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS users;