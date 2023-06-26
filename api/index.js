import express from "express";

// for the env files
import * as env from "dotenv";

// importing the router
import appRouter from "./routes/appRouter.js";

// configuring the env
env.config();

// creating a express app
const app = express();

// setting the port
const PORT = process.env.PORT || 3000;

// Running and listening to the server request
app.listen(PORT, () => {
  console.log(`Server running at ${PORT}`);
});

// using the router
app.use(`api/v1`, appRouter);
