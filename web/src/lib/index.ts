// place files you want to import through the `$lib` alias in this folder.

export function nullStr(s?: string): string | undefined {
	return !s ? undefined : s;
}

export function nullDate(d?: Date | string): Date | undefined {
	return !d ? undefined : new Date(d);
}
