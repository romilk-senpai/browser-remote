declare module '*.css';
declare module '*.svg';
declare module '*.png';
declare module 'jquery';

interface Window {
    lastRightClickedElement: HTMLElement | null;
}