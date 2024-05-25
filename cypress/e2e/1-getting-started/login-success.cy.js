describe("Login Form", () => {
  it("should allow a user to type a username and password", () => {
    cy.visit("http://localhost:8080/login");

    cy.intercept("POST", "http://localhost:8080/login").as("loginRequest");

    cy.get('input[name="username"]').type("test");

    cy.get('input[name="password"]').type("test");

    cy.get("form").submit();

    cy.wait("@loginRequest").then((interception) => {
      assert.equal(
        interception.response.statusCode,
        302,
        "Form submitted successfully with 302 status"
      );
    });
  });
});
