import { Router } from "express";
import { addInformation, getInformations } from "../handlers/informations";

const router = Router();

router.get("/", getInformations);

router.post("/", addInformation);

export default router;
