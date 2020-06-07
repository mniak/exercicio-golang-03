CREATE TABLE Link(
    id SERIAL PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    link TEXT NOT NULL
);

INSERT INTO Link (name, link) VALUES
('Camiseta Branca Masculina P', 'https://example.com/camiseta-br-masc-p'),
('Camiseta Branca Masculina M', 'https://example.com/camiseta-br-masc-m'),
('Camiseta Branca Masculina G', 'https://example.com/camiseta-br-masc-g');