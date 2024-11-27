import { Request, Response } from "express";
import { AddInformationDto } from "../dtos/AddInformation";

export function getInformations(request: Request, response: Response) {
  response.send([]);
}

export function addInformation(request: Request<{}, {}, AddInformationDto>, response: Response) {

}
