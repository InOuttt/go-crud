INSERT INTO user (id, name, password, salt, created_at, updated_at)
VALUES ('967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f14', 'user1', '', '', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp),
       ('c809bf15-bc2c-4621-bb96-70af96fd5d62', 'user2', '', '', '2019-10-02 11:16:12'::timestamp, '2019-10-02 11:16:12'::timestamp),
       ('2367710a-d4fb-49f5-8860-557b337386dr', 'user3', '', '', '2019-10-05 05:21:11'::timestamp, '2019-10-05 05:21:11'::timestamp),
       ('b0a24f12-428f-4ff5-84d5-bc1fdcff6f01', 'user4', '', '', '2019-10-11 19:43:18'::timestamp, '2019-10-11 19:43:18'::timestamp),
       ('e0bb80ec-75a6-4348-bfc3-6ac1e89b195w', 'user5', '', '', '2019-10-12 12:16:02'::timestamp, '2019-10-12 12:16:02'::timestamp);

INSERT INTO profile (id, userId, email, address, created_at, updated_at)
VALUES ('769d5bb5-3a7a-4d5e-8a6c-70af96fd5d66', '967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f14', 'user1@mail.com', 'address for user 1', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp),
       ('80c9bf15-bc2c-4621-bb96-febc8c5b3f11', 'c809bf15-bc2c-4621-bb96-70af96fd5d62', 'user2@mail.com', 'address for user 2', '2019-10-02 11:16:12'::timestamp, '2019-10-02 11:16:12'::timestamp),
       ('3627710a-d4fb-49f5-8860-bc1fdcff6f09', '2367710a-d4fb-49f5-8860-557b337386dr', 'user3@mail.com', 'address for user 3', '2019-10-05 05:21:11'::timestamp, '2019-10-05 05:21:11'::timestamp),
       ('0ba24f12-428f-4ff5-84d5-557b337386rd', 'b0a24f12-428f-4ff5-84d5-bc1fdcff6f01', 'user4@mail.com', 'address for user 4', '2019-10-11 19:43:18'::timestamp, '2019-10-11 19:43:18'::timestamp),
       ('b0eb80ec-75a6-4348-bfc3-95w1e89b16ac', 'e0bb80ec-75a6-4348-bfc3-6ac1e89b195w', 'user5@mail.com', 'address for user 5', '2019-10-12 12:16:02'::timestamp, '2019-10-12 12:16:02'::timestamp);
