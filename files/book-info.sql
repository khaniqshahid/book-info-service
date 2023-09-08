CREATE DATABASE bookdetails;
USE bookdetails;

DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `book_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `author` varchar(100) NOT NULL,
  `publisher` varchar(100) NOT NULL,
  `price` decimal(5,2) NOT NULL,
  `issued_at` date NOT NULL,
  `description` varchar(100) NOT NULL,
  	
  PRIMARY KEY (`book_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2004 DEFAULT CHARSET=latin1;
INSERT INTO `books` VALUES
	(2000,'The Java Programming Language','Shahid','Packt Publishing Ltd','58.99','1988-05-21','A practical guide to the Go programming language.'),
	(2001,'The Python Programming Language','Kushal','Packt Publishing Ltd','58.99','1988-05-21','A practical guide to the Go programming language.'),
	(2002,'The Go Programming Language','Shahid','Packt Publishing Ltd','58.99','1988-05-21','A practical guide to the Go programming language.'),
	(2003,'The C Programming Language','Shahid','Packt Publishing Ltd','58.99','1988-05-21','A practical guide to the Go programming language.');


DROP TABLE IF EXISTS `reviews`;
CREATE TABLE `reviews` (
  `review_id` int(11) NOT NULL AUTO_INCREMENT,
  `book_id` int(11) NOT NULL,
  `review` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`review_id`),
  KEY `reviews_FK` (`book_id`),
  CONSTRAINT `reviews_FK` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;
INSERT INTO `reviews` VALUES
	(95470,2000,1),
	(95471,2001,0),
    (95472,2002,1),
    (95473,2003,4);
