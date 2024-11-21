import express, { Request, Response } from "express";
import dotenv from "dotenv";

dotenv.config();

const app = express();
const PORT = process.env.PORT || 3500;

app.use(express.json());

app.get("/", (req: Request, res: Response) => {
    res.send("Hello, TypeScript with ES Modules!");
});

app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});