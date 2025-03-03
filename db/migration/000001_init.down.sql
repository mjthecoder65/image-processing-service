DROP TABLE IF EXISTS image_transformations;
DROP INDEX IF EXISTS idx_image_transformations_image_id;

DROP TABLE IF EXISTS images;
DROP INDEX IF EXISTS idx_images_user_id;

DROP TRIGGER IF EXISTS update_users_timestamp ON users;
DROP FUNCTION IF EXISTS update_timestamp;

DROP TABLE IF EXISTS users;