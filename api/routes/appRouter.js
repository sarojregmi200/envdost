import express from "express";

// importing the controllers
import { register, login } from "../controllers/_authController.js";

// main entry point for all the routes
const appRouter = express.Router();

// auth routes
appRouter.route(`/auth/register/`).post(register); // logins the user
appRouter.route(`/auth/login/`).post(login); // registers the user
