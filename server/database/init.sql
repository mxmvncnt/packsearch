INSERT INTO public.distro (id, name, version)
VALUES (1, 'Ubuntu', '25.04'),
       (2, 'Ubuntu', '25.10'),
       (3, 'Fedora', '42'),
       (4, 'Fedora', '43');

INSERT INTO public.distro (id, name)
VALUES (5, 'Arch'),
       (6, 'Flathub'),
       (7, 'Snap'),
       (8, 'AppImage'),
       (9, 'Homebrew');

-- Data for Name: distro
INSERT INTO public.package (id, human_name, name, latest_version, description, keywords, homepage, developer)
VALUES (4, 'Java Open JDK 21', 'java_open_jdk_21', 1, 'OpenJDK 21 Runtime Environment', ARRAY['java', 'openjdk', 'jdk21', 'jdk'], 'https://openjdk.org/projects/jdk/21/', ARRAY['Open JDK']),
       (1, 'Java Open JDK 17', 'java_open_jdk_17', 2, 'OpenJDK 17 Runtime Environment', ARRAY['java', 'openjdk', 'jdk17', 'jdk'], 'https://openjdk.org/projects/jdk/17/', ARRAY['OpenJDK']),
       (5, 'Java Open JDK', 'open_jdk', 1, 'OpenJDK 17 Runtime Environment', ARRAY['jdk', 'java', 'openjdk'], 'https://openjdk.org/projects/', ARRAY['OpenJDK']),
       (2, 'Firefox', 'firefox', 3, 'Mozilla Firefox Web browser', ARRAY['mozilla'], 'https://www.mozilla.org/fr/firefox/new/', ARRAY['Mozilla']);

-- Data for Name: major_version
INSERT INTO public.major_version (id, package_id, version_name)
VALUES (1, 5, '17'),
       (2, 5, '21'),
       (3, 2, '123');

-- Data for Name: version
INSERT INTO public.version (id, major_version_id, version_name)
VALUES (1, 2, '21.0.2'),
       (2, 1, '17.0.9'),
       (3, 3, '123.0');

-- Data for Name: variation
INSERT INTO public.variation (id, package_id, distro_id, name, version, package_url, download_url)
VALUES (3, 2, 4, 'org.mozilla.firefox', 3, 'https://flathub.org/apps/org.mozilla.firefox', 'https://dl.flathub.org/repo/appstream/org.mozilla.firefox.flatpakref'),
       (1, 1, 2, 'java-17-openjdk', 2, 'https://src.fedoraproject.org/rpms/java-17-openjdk', 'https://kojipkgs.fedoraproject.org//packages/java-17-openjdk/17.0.9.0.9/3.fc39/src/java-17-openjdk-17.0.9.0.9-3.fc39.src.rpm'),
       (2, 1, 1, 'openjdk-17-jdk', 2, 'https://packages.ubuntu.com/focal/openjdk-17-jdk', 'http://security.ubuntu.com/ubuntu/pool/universe/o/openjdk-17/openjdk-17-jdk_17.0.9+9-1~20.04_amd64.deb');

INSERT INTO public.variation (id, package_id, distro_id, name, version, package_url)
VALUES (4, 1, 9, 'openjdk@17', 2, 'https://formulae.brew.sh/formula/openjdk@17'),
       (7, 4, 9, 'openjdk@21', 1, 'https://formulae.brew.sh/formula/openjdk');-- Data for Name: package
