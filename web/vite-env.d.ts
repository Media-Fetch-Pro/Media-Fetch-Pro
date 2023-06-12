  declare module 'element-plus' {
  export class ElMessage {
    static success(message: string): void;
    static warning(message: string): void;
    static info(message: string): void;
    static error(message: string): void;
  }
}
