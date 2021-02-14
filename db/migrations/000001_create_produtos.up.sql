CREATE TABLE `Produtos` (
	`Id` INT NOT NULL AUTO_INCREMENT,
	`IdRestaurante` INT NOT NULL,
	`Foto` blob NOT NULL,
	`Nome` varchar(255) NOT NULL,
	`Preço` INT NOT NULL,
	`Categoria` varchar(255) NOT NULL,
	`Promoção` BOOLEAN NOT NULL,
	`DescriçãoPromoção` varchar(255) NOT NULL,
	`PreçoPromoção` INT NOT NULL,
	PRIMARY KEY (`Id`)
);

CREATE TABLE `Restaurante` (
	`Id` INT NOT NULL AUTO_INCREMENT,
	`Foto` blob NOT NULL,
	`Nome` varchar(255) NOT NULL,
	`Endereço` varchar(255) NOT NULL,
	`Funcionamento` varchar(255) NOT NULL,
	PRIMARY KEY (`Id`)
);

CREATE TABLE `Promoção` (
	`Id` INT NOT NULL AUTO_INCREMENT,
	`IdProduto` INT NOT NULL,
	`DiaDaSemana` INT NOT NULL,
	`HorarioInicio` varchar(5) NOT NULL,
	`HorarioFim` varchar(5) NOT NULL,
	PRIMARY KEY (`Id`)
);

ALTER TABLE `Produtos` ADD CONSTRAINT `Produtos_fk0` FOREIGN KEY (`IdRestaurante`) REFERENCES `Restaurante`(`Id`);

ALTER TABLE `Promoção` ADD CONSTRAINT `Promoção_fk0` FOREIGN KEY (`IdProduto`) REFERENCES `Produtos`(`Id`);

