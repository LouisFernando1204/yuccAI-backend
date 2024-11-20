import express, { Request, Response } from "express";
import dotenv from 'dotenv';

dotenv.config();

const app = express();
const path = require('path');
const PORT = process.env.PORT || 3500;

app.use(express.urlencoded({ extended: false }));
app.use(express.json());

app.get('/', (req: Request, res: Response) => {
    res.send('Hello Yobel!');
});

app.listen(PORT, () => console.log(`Server running on port ${PORT}`));