import {describe, expect, test} from '@jest/globals';
import { getCleanUrl } from './bilibili'

describe('sum getCleanUrl', () => {
    test('clear login bilibili url', () => {
        expect(getCleanUrl("https://www.bilibili.com/video/BV1cX4y177WM/?spm_id_from=333.1007.tianma.1-2-2.click&vd_source=5d78b061956227324f710505619f52d7")).toBe("https://www.bilibili.com/video/BV1cX4y177WM/");
        expect(getCleanUrl("https://www.bilibili.com/video/BV1p14y167fY/?spm_id_from=333.1007.tianma.1-2-2.click&vd_source=5d78b061956227324f710505619f52d7")).toBe("https://www.bilibili.com/video/BV1p14y167fY/");
        expect(getCleanUrl("https://www.bilibili.com/video/BV1pP41167zy/?spm_id_from=333.1007.tianma.1-1-1.click")).toBe("https://www.bilibili.com/video/BV1pP41167zy/")
        expect(getCleanUrl("https://www.bilibili.com/video/BV1Zp4y1G7MP/?spm_id_from=333.788.recommend_more_video.-1")).toBe("https://www.bilibili.com/video/BV1Zp4y1G7MP/")
    });
});

