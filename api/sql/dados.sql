insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);
 
insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Teste publicação do usuário 1! Deu certo!", 1),
("Publicação do Usuário 2", "Teste publicação do usuário 2! Deu certo!", 2),
("Publicação do Usuário 3", "Teste publicação do usuário 3! Deu certo!", 3);