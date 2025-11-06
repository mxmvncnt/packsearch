-- Data for Name: distro
INSERT INTO public.distro (id, name, version)
VALUES (1, 'Ubuntu', '25.04'),
       (2, 'Ubuntu', '25.10'),
       (3, 'Fedora', '42'),
       (4, 'Fedora', '43'),
       (5, 'Arch', NULL),
       (6, 'Flathub', NULL),
       (7, 'Snap', NULL),
       (8, 'AppImage', NULL),
       (9, 'Homebrew', NULL);

-- Data for Name: package
INSERT INTO public.package (id, human_name, name, latest_version, description, keywords, homepage, developer)
VALUES (3, 'All null values where possible', NULL, NULL, NULL, NULL, NULL, ARRAY['me']),
       (4, 'Java Open JDK 21', 'java_open_jdk_21', NULL, 'OpenJDK 21 Runtime Environment', ARRAY['java', 'openjdk',
           'jdk21', 'jdk'], 'https://openjdk.org/projects/jdk/21/', ARRAY['Open JDK']),
       (1, 'Java Open JDK 17', 'java_open_jdk_17', NULL, 'OpenJDK 17 Runtime Environment', ARRAY['java', 'openjdk',
           'jdk17', 'jdk'], 'https://openjdk.org/projects/jdk/17/', ARRAY['OpenJDK']),
       (5, 'Java Open JDK', 'open_jdk', 1, 'OpenJDK 17 Runtime Environment', ARRAY['jdk', 'java', 'openjdk'],
        'https://openjdk.org/projects/', ARRAY['OpenJDK']),
       (2, 'Firefox', 'firefox', 3, 'Mozilla Firefox Web browser', ARRAY['mozilla'],
        'https://www.mozilla.org/fr/firefox/new/', ARRAY['Mozilla']);

-- Data for Name: major_version
INSERT INTO public.major_version (id, package_id, version_name, release_date)
VALUES (1, 5, '17', NULL),
       (2, 5, '21', NULL),
       (3, 2, '123', NULL);

-- Data for Name: version
INSERT INTO public.version (id, major_version_id, version_name, release_date)
VALUES (1, 2, '21.0.2', NULL),
       (2, 1, '17.0.9', NULL),
       (3, 3, '123.0', NULL);

-- Data for Name: variation
INSERT INTO public.variation (id, package_id, distro_id, name, version, package_url, download_url)
VALUES (6, 3, 3, 'java_all_null_values_where_possible', NULL, NULL, NULL),
       (3, 2, 4, 'org.mozilla.firefox', 3, 'https://flathub.org/apps/org.mozilla.firefox',
        'https://dl.flathub.org/repo/appstream/org.mozilla.firefox.flatpakref'),
       (4, 1, 9, 'openjdk@17', 2, 'https://formulae.brew.sh/formula/openjdk@17', NULL),
       (1, 1, 2, 'java-17-openjdk', 2, 'https://src.fedoraproject.org/rpms/java-17-openjdk',
        'https://kojipkgs.fedoraproject.org//packages/java-17-openjdk/17.0.9.0.9/3.fc39/src/java-17-openjdk-17.0.9.0.9-3.fc39.src.rpm'),
       (7, 4, 9, 'openjdk@21', 1, 'https://formulae.brew.sh/formula/openjdk', NULL),
       (2, 1, 1, 'openjdk-17-jdk', 2, 'https://packages.ubuntu.com/focal/openjdk-17-jdk',
        'http://security.ubuntu.com/ubuntu/pool/universe/o/openjdk-17/openjdk-17-jdk_17.0.9+9-1~20.04_amd64.deb');
