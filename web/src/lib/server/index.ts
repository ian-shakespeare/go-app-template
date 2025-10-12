import createClient from "openapi-fetch"
import type { paths } from "../schema"
import { API_URL } from "$env/static/private"

export const apiClient = createClient<paths>({ baseUrl: API_URL });
