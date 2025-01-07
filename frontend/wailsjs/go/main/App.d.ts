// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {backend} from '../models';

export function ExportPdfs(arg1:Array<string>):Promise<void>;

export function GetCheckedFiles():Promise<Array<backend.DocInfo>>;

export function GetCheckedFilesCount():Promise<number>;

export function GetElementsByIds(arg1:Array<string>):Promise<Array<backend.DocInfo>>;

export function GetPaths(arg1:Array<string>):Promise<Array<string>>;

export function GetTabletFolder(arg1:string):Promise<Array<backend.DocInfo>>;

export function GetTabletFolderSelection(arg1:string):Promise<Array<backend.SelectionInfo>>;

export function IsIpValid(arg1:string):Promise<boolean>;

export function OnItemSelect(arg1:string,arg2:boolean):Promise<void>;

export function ReadTabletDocs(arg1:string):Promise<void>;
