// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {backend} from '../models';

export function DirectoryDialog():Promise<string>;

export function Export():Promise<void>;

export function GetAppVersion():Promise<string>;

export function GetCheckedFiles():Promise<Array<backend.DocInfo>>;

export function GetCheckedFilesCount():Promise<number>;

export function GetExportOptions():Promise<backend.RmExportOptions>;

export function GetFolder(arg1:string):Promise<Array<backend.DocInfo>>;

export function GetFolderSelection(arg1:string):Promise<Array<backend.SelectionInfo>>;

export function GetItemSelection(arg1:string):Promise<backend.SelectionInfo>;

export function InitExport():Promise<void>;

export function IsIpValid(arg1:string):Promise<boolean>;

export function OnItemSelect(arg1:string,arg2:boolean):Promise<void>;

export function ReadDocs(arg1:string):Promise<void>;

export function SetExportOptions(arg1:backend.RmExportOptions):Promise<void>;
