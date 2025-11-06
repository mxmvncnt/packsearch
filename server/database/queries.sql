-- name: GetPackageFromID :one
SELECT *
FROM package
WHERE id = @id::numeric;

-- name: GetAllPackages :many
SELECT id,
    human_name,
    name,
    latest_version,
    description,
    keywords,
    homepage,
    developer
FROM package;

-- name: FuzzySearch :many
SELECT DISTINCT package.id,
                human_name,
                package.name,
                latest_version,
                description,
                keywords,
                homepage,
                developer
FROM package
         INNER JOIN public.variation v ON package.id = v.package_id
WHERE dmetaphone(@term::text) ILIKE dmetaphone(human_name)
   OR dmetaphone(@term::text) = ANY (ARRAY(SELECT dmetaphone(element) FROM unnest(keywords) AS element))
   OR dmetaphone(@term::text) = ANY (ARRAY(SELECT dmetaphone(element) FROM unnest(developer) AS element))
   OR dmetaphone(@term::text) ILIKE dmetaphone(human_name)
   OR dmetaphone(@term::text) ILIKE dmetaphone(package.name)
   OR dmetaphone(@term::text) ILIKE dmetaphone(description)
   OR dmetaphone(@term::text) ILIKE dmetaphone(v.name);

-- name: GetVariations :many
SELECT variation.id      AS id,
       package_id,
       distro_id,
       p.human_name      AS human_name,
       variation.name    AS name,
       variation.version AS version,
       package_url,
       download_url,
       d.name            AS distro_name,
       d.version         AS distro_version
FROM variation
         INNER JOIN distro d ON d.id = variation.distro_id
         INNER JOIN package p ON p.id = variation.package_id
WHERE package_id = @package_id::numeric;