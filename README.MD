# State

-   ## Working
    -   Sign-Up
    -   Sign-In
    -   Create/Update/Delete Posts
    -   Create/Update/Delete Comments
    -   Routes should be protected on both end
-   # Todo
    -   Error handling/feedbach (Front & Back)
    -   Logging/Tracing (Back)
    -   Refactoring (mostly Front)
    -   Docs (Back)
    -   Migrations (Back)
    -   Performance (mostly Front)
-   # How to run

    -   ## Front
        -   <code>npm i && npm run dev</code> which would make app run on port <b>3000</b>
    -   ## Back
        -   <code>docker-compose -f .\deployment\docker-compose.yml up -d --force-recreate</code> would start the postgres/adminer (<i>adminer is UI to manage database, port specified for it is <b>3333</b></i>)
        -   env file is included and can be modifed in <code>configs/.env</code>
        -   <code>db.sql</code> contains schema that can be easly imported with <code>adminer on port 3333</code>
        -   <code>go run /cmd</code> should be sufficient enough to run the app

-   # Notes
    -   wouldn't it make sense to split it into microservices and make a use of NoSQL?
