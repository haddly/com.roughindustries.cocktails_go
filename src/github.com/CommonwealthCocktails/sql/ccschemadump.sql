/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `altIngredient` (
  `idAltIngredient` int(11) NOT NULL AUTO_INCREMENT,
  `idProduct` int(11) NOT NULL,
  `idRecipeStep` int(11) NOT NULL,
  PRIMARY KEY (`idAltIngredient`),
  KEY `altIngredient_idProduct_id_fk` (`idProduct`),
  KEY `altIngredient_idRecipeStep_id_fk` (`idRecipeStep`),
  CONSTRAINT `altIngredient_idProduct_id_fk` FOREIGN KEY (`idProduct`) REFERENCES `product` (`idProduct`),
  CONSTRAINT `altIngredient_idRecipeStep_id_fk` FOREIGN KEY (`idRecipeStep`) REFERENCES `recipestep` (`idRecipeStep`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `altnames` (
  `idAltNames` int(11) NOT NULL AUTO_INCREMENT,
  `altNamesString` varchar(250) NOT NULL,
  PRIMARY KEY (`idAltNames`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cocktail` (
  `idCocktail` int(11) NOT NULL AUTO_INCREMENT,
  `cocktailTitle` varchar(150) NOT NULL,
  `cocktailName` varchar(150) NOT NULL,
  `cocktailDisplayName` varchar(150) DEFAULT NULL,
  `cocktailSpokenName` varchar(150) DEFAULT NULL,
  `cocktailOrigin` varchar(2500) DEFAULT NULL,
  `cocktailDescription` varchar(2500) DEFAULT NULL,
  `cocktailComment` varchar(2500) DEFAULT NULL,
  `cocktailImagePath` varchar(1000) DEFAULT NULL,
  `cocktailImage` varchar(250) DEFAULT NULL,
  `cocktailImageSourceName` varchar(250) DEFAULT NULL,
  `cocktailImageSourceLink` varchar(1000) DEFAULT NULL,
  `cocktailRating` int(1) DEFAULT NULL,
  `cocktailSourceName` varchar(150) DEFAULT NULL,
  `cocktailSourceLink` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`idCocktail`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cocktailToAKAs` (
  `idCocktailToAKANames` int(11) NOT NULL AUTO_INCREMENT,
  `idCocktail` int(11) NOT NULL,
  `idAKAName` int(11) NOT NULL,
  PRIMARY KEY (`idCocktailToAKANames`),
  KEY `cocktailToAKAs_idCocktail_id_fk` (`idCocktail`),
  KEY `cocktailToAKAs_idAKAName_id_fk` (`idAKAName`),
  CONSTRAINT `cocktailToAKAs_idAKAName_id_fk` FOREIGN KEY (`idAKAName`) REFERENCES `altnames` (`idAltNames`),
  CONSTRAINT `cocktailToAKAs_idCocktail_id_fk` FOREIGN KEY (`idCocktail`) REFERENCES `cocktail` (`idCocktail`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cocktailToAltNames` (
  `idCocktailToAltNames` int(11) NOT NULL AUTO_INCREMENT,
  `idCocktail` int(11) NOT NULL,
  `idAltName` int(11) NOT NULL,
  PRIMARY KEY (`idCocktailToAltNames`),
  KEY `cocktailToAltNames_idCocktail_id_fk` (`idCocktail`),
  KEY `cocktailToAltNames_idAltName_id_fk` (`idAltName`),
  CONSTRAINT `cocktailToAltNames_idAltName_id_fk` FOREIGN KEY (`idAltName`) REFERENCES `altnames` (`idAltNames`),
  CONSTRAINT `cocktailToAltNames_idCocktail_id_fk` FOREIGN KEY (`idCocktail`) REFERENCES `cocktail` (`idCocktail`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cocktailToMetas` (
  `idCocktailToMetas` int(11) NOT NULL AUTO_INCREMENT,
  `idCocktail` int(11) NOT NULL,
  `idMeta` int(11) NOT NULL,
  `idMetaType` int(11) NOT NULL,
  `isRootCocktailForMeta` tinyint(1) NOT NULL,
  PRIMARY KEY (`idCocktailToMetas`),
  KEY `cocktailToMetas_idCocktail_id_fk` (`idCocktail`),
  KEY `cocktailToMetas_idMeta_id_fk` (`idMeta`),
  KEY `cocktailToMetas_idMetaType_id_fk` (`idMetaType`),
  CONSTRAINT `cocktailToMetas_idCocktail_id_fk` FOREIGN KEY (`idCocktail`) REFERENCES `cocktail` (`idCocktail`),
  CONSTRAINT `cocktailToMetas_idMetaType_id_fk` FOREIGN KEY (`idMetaType`) REFERENCES `metatype` (`idMetaType`),
  CONSTRAINT `cocktailToMetas_idMeta_id_fk` FOREIGN KEY (`idMeta`) REFERENCES `meta` (`idMeta`)
) ENGINE=InnoDB AUTO_INCREMENT=313 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cocktailToPosts` (
  `idCocktailToPosts` int(11) NOT NULL AUTO_INCREMENT,
  `idCocktail` int(11) NOT NULL,
  `idPost` int(11) NOT NULL,
  PRIMARY KEY (`idCocktailToPosts`),
  KEY `cocktailToPosts_idCocktail_id_fk` (`idCocktail`),
  KEY `cocktailToPosts_idPost_id_fk` (`idPost`),
  CONSTRAINT `cocktailToPosts_idCocktail_id_fk` FOREIGN KEY (`idCocktail`) REFERENCES `cocktail` (`idCocktail`),
  CONSTRAINT `cocktailToPosts_idPost_id_fk` FOREIGN KEY (`idPost`) REFERENCES `post` (`idPost`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cocktailToProducts` (
  `idCocktailToProducts` int(11) NOT NULL AUTO_INCREMENT,
  `idCocktail` int(11) NOT NULL,
  `idProduct` int(11) NOT NULL,
  `idProductType` int(11) NOT NULL,
  PRIMARY KEY (`idCocktailToProducts`),
  KEY `cocktailToProducts_idCocktail_id_fk` (`idCocktail`),
  KEY `cocktailToProducts_idProduct_id_fk` (`idProduct`),
  KEY `cocktailToProducts_idProductType_id_fk` (`idProductType`),
  CONSTRAINT `cocktailToProducts_idCocktail_id_fk` FOREIGN KEY (`idCocktail`) REFERENCES `cocktail` (`idCocktail`),
  CONSTRAINT `cocktailToProducts_idProductType_id_fk` FOREIGN KEY (`idProductType`) REFERENCES `producttype` (`idProductType`),
  CONSTRAINT `cocktailToProducts_idProduct_id_fk` FOREIGN KEY (`idProduct`) REFERENCES `product` (`idProduct`)
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cocktailToRecipe` (
  `idCocktailToRecipe` int(11) NOT NULL AUTO_INCREMENT,
  `idCocktail` int(11) NOT NULL,
  `idRecipe` int(11) NOT NULL,
  PRIMARY KEY (`idCocktailToRecipe`),
  KEY `cocktailToRecipe_idCocktail_id_fk` (`idCocktail`),
  KEY `cocktailToRecipe_idRecipe_id_fk` (`idRecipe`),
  CONSTRAINT `cocktailToRecipe_idCocktail_id_fk` FOREIGN KEY (`idCocktail`) REFERENCES `cocktail` (`idCocktail`),
  CONSTRAINT `cocktailToRecipe_idRecipe_id_fk` FOREIGN KEY (`idRecipe`) REFERENCES `recipe` (`idRecipe`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `derivedProduct` (
  `idDerivedProduct` int(11) NOT NULL AUTO_INCREMENT,
  `idBaseProduct` int(11) NOT NULL,
  `idProduct` int(11) NOT NULL,
  PRIMARY KEY (`idDerivedProduct`),
  KEY `derivedProduct_idBaseProduct_id_fk` (`idBaseProduct`),
  KEY `derivedProduct_idProduct_id_fk` (`idProduct`),
  CONSTRAINT `derivedProduct_idBaseProduct_id_fk` FOREIGN KEY (`idBaseProduct`) REFERENCES `product` (`idProduct`),
  CONSTRAINT `derivedProduct_idProduct_id_fk` FOREIGN KEY (`idProduct`) REFERENCES `product` (`idProduct`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `doze` (
  `idDoze` int(11) NOT NULL AUTO_INCREMENT,
  `dozeName` varchar(150) NOT NULL,
  PRIMARY KEY (`idDoze`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `groupProduct` (
  `idGroupProduct` int(11) NOT NULL AUTO_INCREMENT,
  `idBaseProduct` int(11) NOT NULL,
  `idProduct` int(11) NOT NULL,
  PRIMARY KEY (`idGroupProduct`),
  KEY `groupProduct_idProduct_id_fk` (`idProduct`),
  CONSTRAINT `groupProduct_idGroup_id_fk` FOREIGN KEY (`idGroupProduct`) REFERENCES `product` (`idProduct`),
  CONSTRAINT `groupProduct_idProduct_id_fk` FOREIGN KEY (`idProduct`) REFERENCES `product` (`idProduct`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `grouptype` (
  `idGroupType` int(11) NOT NULL AUTO_INCREMENT,
  `groupTypeName` varchar(150) NOT NULL,
  PRIMARY KEY (`idGroupType`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meta` (
  `idMeta` int(11) NOT NULL AUTO_INCREMENT,
  `metaName` varchar(150) NOT NULL,
  `metaType` int(11) NOT NULL,
  `metaArticle` int(11) DEFAULT NULL,
  `metaBlurb` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`idMeta`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `metatype` (
  `idMetaType` int(11) NOT NULL AUTO_INCREMENT,
  `metatypeShowInCocktailsIndex` tinyint(1) DEFAULT NULL,
  `metatypeName` varchar(150) NOT NULL,
  `metatypeOrdinal` int(11) DEFAULT NULL,
  `metatypeHasRoot` tinyint(1) DEFAULT NULL,
  `metatypeIsOneToMany` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`idMetaType`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product` (
  `idProduct` int(11) NOT NULL AUTO_INCREMENT,
  `productName` varchar(150) NOT NULL,
  `productType` int(11) NOT NULL,
  `productDescription` varchar(1500) DEFAULT NULL,
  `productDetails` varchar(1500) DEFAULT NULL,
  `productImagePath` varchar(750) DEFAULT NULL,
  `productImage` varchar(500) DEFAULT NULL,
  `productImageSourceName` varchar(500) DEFAULT NULL,
  `productImageSourceLink` varchar(750) DEFAULT NULL,
  `productArticle` int(11) DEFAULT NULL,
  `productRecipe` int(11) DEFAULT NULL,
  `productGroupType` int(11) DEFAULT NULL,
  `productPreText` varchar(250) DEFAULT NULL,
  `productPostText` varchar(250) DEFAULT NULL,
  `productRating` int(1) DEFAULT NULL,
  `productSourceName` varchar(1500) DEFAULT NULL,
  `productSourceLink` varchar(1500) DEFAULT NULL,
  `productAbout` int(11) DEFAULT NULL,
  `productAmazonLink` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`idProduct`),
  KEY `product_producttype_id_fk` (`productType`),
  KEY `product_productgrouptype_id_fk` (`productGroupType`),
  KEY `product_productarticle_id_fk` (`productArticle`),
  KEY `product_productrecipe_id_fk` (`productRecipe`),
  KEY `product_productabout_id_fk` (`productAbout`),
  CONSTRAINT `product_productabout_id_fk` FOREIGN KEY (`productAbout`) REFERENCES `post` (`idPost`),
  CONSTRAINT `product_productarticle_id_fk` FOREIGN KEY (`productArticle`) REFERENCES `post` (`idPost`),
  CONSTRAINT `product_productgrouptype_id_fk` FOREIGN KEY (`productGroupType`) REFERENCES `grouptype` (`idGroupType`),
  CONSTRAINT `product_productrecipe_id_fk` FOREIGN KEY (`productRecipe`) REFERENCES `recipe` (`idRecipe`),
  CONSTRAINT `product_producttype_id_fk` FOREIGN KEY (`productType`) REFERENCES `producttype` (`idProductType`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `producttype` (
  `idProductType` int(11) NOT NULL AUTO_INCREMENT,
  `productTypeName` varchar(150) NOT NULL,
  `productTypeIsIngredient` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`idProductType`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recipe` (
  `idRecipe` int(11) NOT NULL AUTO_INCREMENT,
  `recipeMethod` varchar(2500) DEFAULT NULL,
  PRIMARY KEY (`idRecipe`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recipeToRecipeSteps` (
  `idRecipeToRecipeSteps` int(11) NOT NULL AUTO_INCREMENT,
  `idRecipe` int(11) NOT NULL,
  `idRecipeStep` int(11) NOT NULL,
  PRIMARY KEY (`idRecipeToRecipeSteps`),
  KEY `recipeToRecipeSteps_idRecipe_id_fk` (`idRecipe`),
  KEY `recipeToRecipeSteps_idRecipeStep_id_fk` (`idRecipeStep`),
  CONSTRAINT `recipeToRecipeSteps_idRecipeStep_id_fk` FOREIGN KEY (`idRecipeStep`) REFERENCES `recipestep` (`idRecipeStep`),
  CONSTRAINT `recipeToRecipeSteps_idRecipe_id_fk` FOREIGN KEY (`idRecipe`) REFERENCES `recipe` (`idRecipe`)
) ENGINE=InnoDB AUTO_INCREMENT=189 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recipestep` (
  `idRecipeStep` int(11) NOT NULL AUTO_INCREMENT,
  `recipestepOriginalIngredient` int(11) NOT NULL,
  `recipestepRecipeCardinalFloat` float NOT NULL,
  `recipestepRecipeCardinalString` varchar(15) NOT NULL,
  `recipestepRecipeOrdinal` int(11) NOT NULL,
  `recipestepRecipeDoze` int(11) NOT NULL,
  `recipestepAdIngredient` int(11) DEFAULT NULL,
  PRIMARY KEY (`idRecipeStep`),
  KEY `recipestep_recipestepadingredient_id_fk` (`recipestepAdIngredient`),
  KEY `recipestep_recipesteprecipedoze_id_fk` (`recipestepRecipeDoze`),
  KEY `recipestep_recipesteporiginalingredient_id_fk` (`recipestepOriginalIngredient`),
  CONSTRAINT `recipestep_recipestepadingredient_id_fk` FOREIGN KEY (`recipestepAdIngredient`) REFERENCES `product` (`idProduct`),
  CONSTRAINT `recipestep_recipesteporiginalingredient_id_fk` FOREIGN KEY (`recipestepOriginalIngredient`) REFERENCES `product` (`idProduct`),
  CONSTRAINT `recipestep_recipesteprecipedoze_id_fk` FOREIGN KEY (`recipestepRecipeDoze`) REFERENCES `doze` (`idDoze`)
) ENGINE=InnoDB AUTO_INCREMENT=191 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `idUser` int(11) NOT NULL AUTO_INCREMENT,
  `userName` varchar(150) NOT NULL,
  `userPassword` varchar(250) NOT NULL,
  `userEmail` varchar(250) DEFAULT NULL,
  `userLastLogin` datetime NOT NULL,
  PRIMARY KEY (`idUser`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usersessions` (
  `idSession` int(11) NOT NULL AUTO_INCREMENT,
  `idUser` int(11) DEFAULT NULL,
  `usersessionSessionKey` varchar(150) NOT NULL,
  `usersessionCSRFGenTime` datetime NOT NULL,
  `usersessionCSRFBase` varchar(150) NOT NULL,
  `usersessionLastSeenTime` datetime NOT NULL,
  `usersessionCSRFKey` blob NOT NULL,
  `usersessionLoginTime` datetime NOT NULL,
  `usersessionIsDefaultUser` tinyint(1) NOT NULL,
  PRIMARY KEY (`idSession`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
