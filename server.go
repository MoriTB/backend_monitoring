package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Domain struct {
	Domain_name string `json:"domain"`
}
type user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type jwtCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

type address struct {
	Adress string `json:"domain"`
}

var db *sql.DB
var serverKey string = "server-secret-key"

func main() {
	now := time.Now()
	fmt.Println(now.Format("2022-01-02"))
	// create the connection based on the database that has been created on the localhost.
	connStr := "postgresql://postgres:1@localhost/database?sslmode=disable"
	
	var connectionError error
	// opening database and error checking process.
	db, connectionError = sql.Open("postgres", connStr)
	// check the procedure of failing or connecting.
	if connectionError != nil {
		log.Fatal(connectionError)
	} else {
		fmt.Println("connected")
	}

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(serverKey),
	}
	// close database

	restApp := echo.New()
	restApp.GET("/", HelloWorld)
	restApp.POST("/addUser", addUser)
	restApp.GET("/getToken", newToken)
	//end point with middleware to make sure that token is valid
	restApp.POST("/newDomain", newAddress, echojwt.WithConfig(config))
	restApp.GET("/getUserDomains", getDomains, echojwt.WithConfig(config))
	//visit the specefied domain in the request body and set it's visit = vist+1
	restApp.GET("/viewDomain", addViewToDomain)
	restApp.GET("/getDomainWarning", getDomainWarning, echojwt.WithConfig(config))
	//start the server
	restApp.Logger.Fatal(restApp.Start(":1323"))
}

func HelloWorld(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!!!")

}

func addUser(c echo.Context) error {
	u := new(user)
	c.Bind(u)
	// using $1, $2, and $3 as placeholders for the u.Username, u.Email, and u.Password values.
	// those placeholders indicate the position of the argument.
	// to insert data into a "users" table in our PostgreSQL database.
	_, err := db.Exec("INSERT INTO users (username, email,password) VALUES ($1, $2,$3)", u.Username, u.Email, u.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("User added")
	return c.String(http.StatusOK, "userCreatedSucessfully")
}
// echo context has request and response data within.
func newToken(c echo.Context) error {
	// get the username from the request
	// create a new JWT token

	// set claims
	u := new(user)
	c.Bind(u)
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 and password = $2",
		u.Username, u.Password).Scan(&count)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// check if the username exists
	if count == 0 {
		return c.String(http.StatusUnauthorized, "user not found!")
	}

	jwtClaims := jwt.MapClaims{
		"username": u.Username,
	}
	//signing method is HS265
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("server-secret-key"))
	if err != nil {
		return err
	}
	// return the token as part of the response
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func newAddress(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(*jwtCustomClaims)
	username := claims.Username
	domain := new(address)
	c.Bind(domain)
	var domainCount int
	// it checks the domain count with the username registered by token claimed.
	err := db.QueryRow("SELECT COUNT(*) FROM domains WHERE username = $1", username).Scan(&domainCount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if domainCount >= 20 {

		return c.String(http.StatusInternalServerError, "User has reached the maximum number of domains allowed")
	}
	_, err = db.Exec("INSERT INTO domains (username, domain) VALUES ($1, $2)", username, domain.Adress)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("domian %s added to user %s domains", domain.Adress, username))

}

func getDomains(c echo.Context) error {
	jwtToken := c.Get("user").(*jwt.Token)
	claims := jwtToken.Claims.(*jwtCustomClaims)
	username := claims.Username
	rows, err := db.Query("SELECT domain FROM domains WHERE username = $1", username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error querying the domains table",
		})
	}
	defer rows.Close()

	// create a slice to hold the domains
	domains := []Domain{}

	// scan the rows
	for rows.Next() {
		var domain Domain
		if err := rows.Scan(&domain.Domain_name); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Error scanning the rows",
			})
		}

		domains = append(domains, domain)
	}

	// check for errors
	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})

	}
	result := []map[string]string{}
	for _, domain := range domains {
		domainMap := map[string]string{"domain": domain.Domain_name}
		result = append(result, domainMap)
	}
	return c.JSON(http.StatusOK, result)
}

func addViewToDomain(c echo.Context) error {
	domain := new(Domain)
	c.Bind(domain)
	now := time.Now()
	// format date to contain year-month-day
	today := now.Format("2022-01-02")

	// Check if there is already a view for today's date in the views table
	var count int
	err := db.QueryRow("SELECT view_count FROM views WHERE domain = $1 AND view_date = $2", domain.Domain_name, today).Scan(&count)

	if err == sql.ErrNoRows {

		// If there is no view for today's date, insert a new record with view_count = 1
		_, err = db.Exec("INSERT INTO views (domain, view_count, view_date) VALUES ($1, $2, $3)", domain.Domain_name, 1, today)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": err.Error(),
			})
		}
		count = 1
	} else if err != nil {

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	} else {
		count = count + 1
		// If there is already a view for today's date, update the view_count = view_count + 1
		_, err = db.Exec("UPDATE views SET view_count = $1  where domain = $2", count, domain.Domain_name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": err.Error(),
			})
		}

	}
	return c.JSON(http.StatusOK, echo.Map{
		"viewCount": count,
	})

}
func getDomainWarning(c echo.Context) error {
	domain := new(Domain)
	c.Bind(domain)
	now := time.Now()
	today := now.Format("2022-01-02")
	fmt.Println(domain.Domain_name)
	var viewCount int
	err := db.QueryRow("SELECT view_count FROM views WHERE domain = $1 AND view_date = $2", domain.Domain_name, today).Scan(&viewCount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	fmt.Println(viewCount)
	if viewCount > 20 {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": "Domain has reached the warning limit of views today",
		})
	}
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"status": "ok",
	})

}
