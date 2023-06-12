/// <reference types="vite/client" />
import { ElMessage } from "element-plus"
// 手动声明 ElMessage
declare module 'element-plus' {
    export class ElMessage {
        static success(message: string): void;
        static warning(message: string): void;
        static info(message: string): void;
        static error(message: string): void;
    }
}