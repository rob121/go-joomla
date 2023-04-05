# Go-Joomla

A Package for interacting with Joomla

### Features

 * Read joomla configuration
 * Check valid user
 * Check credentials
 * Fetch User from Database
 * Get User Groups

### Examples

```

   //load the joomla config 

    err := joomla.LoadConfig("./configuration.php")

	if err != nil {
		log.Println(err)
	}

	
	//get a configuration item
	out := joomla.Config.Get("helpurl")

	log.Println(out)

	
	//connect to the database
	connd := joomla.Connect()

	if connd != nil {

		log.Println(connd)
		return
	}

    //is this a valid user?

	vu := joomla.ValidUser("admin")

	log.Printf("%+v", vu)


   //Are these valid credentials?

	v2 := joomla.ValidCredentials("admin", "12345")

	log.Printf("%+v", v2)

    //get a joomla user

	v3, _ := joomla.GetUser("admin")

	log.Printf("%+v", v3)

	//Fetch user groups
	
	v4, _ := v3.Groups()

	log.Printf("%+v", v4)


```