/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { PetProperties } from '../models/PetProperties';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class PetService {
    /**
     * Get existing pet
     * @returns PetProperties OK
     * @throws ApiError
     */
    public static getPet(): CancelablePromise<PetProperties> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/pet',
        });
    }
    /**
     * Upload new pet
     * @param petProperties Pet properties
     * @returns any OK
     * @throws ApiError
     */
    public static uploadPet(
        petProperties: PetProperties,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/pet',
            body: petProperties,
            errors: {
                400: `Bad request`,
            },
        });
    }
    /**
     * Delete uploaded pet
     * @returns void
     * @throws ApiError
     */
    public static deletePet(): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/pet',
            errors: {
                400: `Bad request`,
            },
        });
    }
}
