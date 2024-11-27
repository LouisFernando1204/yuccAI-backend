import express, { NextFunction, Request, Response } from "express";
import informationsRouter from './routes/informations'

const app = express();

app.use('/api/informations', informationsRouter)

const PORT = 3000;

app.listen(PORT, () => {
  console.log("Running on port", PORT);
});
